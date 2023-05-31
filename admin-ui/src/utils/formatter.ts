export default {
  articleStatus(row: any) {
    const articleStatus = {
      0: "草稿",
      1: "发布",
    };
    // eslint-disable-next-line camelcase
    const article_Status: keyof typeof articleStatus = row.status;
    return articleStatus[article_Status];
  }

}
