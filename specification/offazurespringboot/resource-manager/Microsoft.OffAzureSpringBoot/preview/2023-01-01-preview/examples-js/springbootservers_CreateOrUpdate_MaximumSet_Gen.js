const { SpringAppDiscoveryManagementClient } = require("@azure/arm-springappdiscovery");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create springbootservers resource.
 *
 * @summary Create springbootservers resource.
 * x-ms-original-file: specification/offazurespringboot/resource-manager/Microsoft.OffAzureSpringBoot/preview/2023-01-01-preview/examples/springbootservers_CreateOrUpdate_MaximumSet_Gen.json
 */
async function springbootserversCreateOrUpdateMaximumSetGen() {
  const subscriptionId = process.env["SPRINGAPPDISCOVERY_SUBSCRIPTION_ID"] || "etmdxomjncqvygm";
  const resourceGroupName =
    process.env["SPRINGAPPDISCOVERY_RESOURCE_GROUP"] || "rgspringbootservers";
  const siteName = "hlkrzldhyobavtabgpubtjbhlslnjmsvkthwcfboriwyxndacjypzbj";
  const springbootserversName = "zkarbqnwnxeozvjrkpdqmgnwedwgtwcmmyqwaijkn";
  const springbootservers = {
    properties: {
      errors: [],
      fqdnAndIpAddressList: [],
      machineArmId: "fvfkiapbqsprnbzczdfmuryknrna",
      port: 10,
      server: "thhuxocfyqpeluqcgnypi",
      springBootApps: 17,
      totalApps: 5,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SpringAppDiscoveryManagementClient(credential, subscriptionId);
  const result = await client.springbootservers.createOrUpdate(
    resourceGroupName,
    siteName,
    springbootserversName,
    springbootservers,
  );
  console.log(result);
}
