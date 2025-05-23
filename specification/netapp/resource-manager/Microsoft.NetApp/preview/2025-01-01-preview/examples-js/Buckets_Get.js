const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get the details of the specified volume's bucket. A bucket allows additional services, such as AI services, connect to the volume data contained in those buckets.
 *
 * @summary Get the details of the specified volume's bucket. A bucket allows additional services, such as AI services, connect to the volume data contained in those buckets.
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/Buckets_Get.json
 */
async function bucketsGet() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "myRG";
  const accountName = "account1";
  const poolName = "pool1";
  const volumeName = "volume1";
  const bucketName = "bucket1";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.buckets.get(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
    bucketName,
  );
  console.log(result);
}
