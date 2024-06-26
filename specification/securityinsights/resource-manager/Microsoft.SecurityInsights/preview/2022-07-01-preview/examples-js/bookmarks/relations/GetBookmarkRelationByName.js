const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a bookmark relation.
 *
 * @summary Gets a bookmark relation.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/bookmarks/relations/GetBookmarkRelationByName.json
 */
async function getABookmarkRelation() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const bookmarkId = "2216d0e1-91e3-4902-89fd-d2df8c535096";
  const relationName = "4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.bookmarkRelations.get(
    resourceGroupName,
    workspaceName,
    bookmarkId,
    relationName
  );
  console.log(result);
}

getABookmarkRelation().catch(console.error);
