package main

const sqlSearch = `
SELECT
  title, subq.workid, act, scene, description, ts_headline(body, q)
FROM (
  SELECT
    workid, act, scene, description, body, ts_rank(tsv, q) as rank, q
  FROM
    scenes, plainto_tsquery($1) q
  WHERE
    tsv @@ q
  ORDER BY rank DESC LIMIT 10
) AS subq
JOIN
  works ON subq.workid = works.workid
ORDER BY
  rank DESC
`

const sqlGetScene = `
SELECT description, body
FROM scenes
WHERE workid=$1 AND act=$2 AND scene=$3
`

const sqlGetWork = `
SELECT title
FROM works
WHERE workid=$1
`
