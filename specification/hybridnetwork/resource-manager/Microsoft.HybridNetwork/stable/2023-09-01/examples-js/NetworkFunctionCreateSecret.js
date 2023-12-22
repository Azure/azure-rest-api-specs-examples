const { HybridNetworkManagementClient } = require("@azure/arm-hybridnetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a network function resource.
 *
 * @summary Creates or updates a network function resource.
 * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionCreateSecret.json
 */
async function createNetworkFunctionResourceWithSecrets() {
  const subscriptionId = process.env["HYBRIDNETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HYBRIDNETWORK_RESOURCE_GROUP"] || "rg";
  const networkFunctionName = "testNf";
  const parameters = {
    location: "eastus",
    properties: {
      allowSoftwareUpdate: false,
      configurationType: "Secret",
      networkFunctionDefinitionVersionResourceReference: {
        id: "/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/testVendor/networkFunctionDefinitionGroups/testnetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.1",
        idType: "Open",
      },
      nfviId:
        "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testResourceGroup/providers/Microsoft.ExtendedLocation/customLocations/testCustomLocation",
      nfviType: "AzureArcKubernetes",
      roleOverrideValues: [
        '{"name":"testRoleOne","deployParametersMappingRuleProfile":{"helmMappingRuleProfile":{"helmPackageVersion":"2.1.3","values":"{\\"roleOneParam\\":\\"roleOneOverrideValue\\"}"}}}',
        '{"name":"testRoleTwo","deployParametersMappingRuleProfile":{"helmMappingRuleProfile":{"releaseName":"overrideReleaseName","releaseNamespace":"overrideNamespace","values":"{\\"roleTwoParam\\":\\"roleTwoOverrideValue\\"}"}}}',
      ],
      secretDeploymentValues: '{"adminPassword":"password1","userPassword":"password2"}',
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridNetworkManagementClient(credential, subscriptionId);
  const result = await client.networkFunctions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    networkFunctionName,
    parameters,
  );
  console.log(result);
}
