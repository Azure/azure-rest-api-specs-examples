const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Datastore
 *
 * @summary Create a Datastore
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/Datastores_CreateOrUpdate.json
 */
async function datastoresCreateOrUpdate() {
  const subscriptionId =
    process.env["AVS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["AVS_RESOURCE_GROUP"] || "group1";
  const privateCloudName = "cloud1";
  const clusterName = "cluster1";
  const datastoreName = "datastore1";
  const datastore = {
    netAppVolume: {
      id: "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/ResourceGroup1/providers/Microsoft.NetApp/netAppAccounts/NetAppAccount1/capacityPools/CapacityPool1/volumes/NFSVol1",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.datastores.beginCreateOrUpdateAndWait(
    resourceGroupName,
    privateCloudName,
    clusterName,
    datastoreName,
    datastore,
  );
  console.log(result);
}
