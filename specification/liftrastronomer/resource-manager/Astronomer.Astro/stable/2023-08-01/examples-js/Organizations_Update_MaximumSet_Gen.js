const { AstroManagementClient } = require("@azure/arm-astro");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a OrganizationResource
 *
 * @summary Update a OrganizationResource
 * x-ms-original-file: specification/liftrastronomer/resource-manager/Astronomer.Astro/stable/2023-08-01/examples/Organizations_Update_MaximumSet_Gen.json
 */
async function organizationsUpdate() {
  const subscriptionId =
    process.env["ASTRO_SUBSCRIPTION_ID"] || "43454B17-172A-40FE-80FA-549EA23D12B3";
  const resourceGroupName = process.env["ASTRO_RESOURCE_GROUP"] || "rgastronomer";
  const organizationName = "6.";
  const properties = {
    identity: { type: "None", userAssignedIdentities: {} },
    properties: {
      partnerOrganizationProperties: {
        organizationId: "lrtmbkvyvvoszhjevohkmyjhfyty",
        organizationName: "U2P_",
        singleSignOnProperties: {
          aadDomains: ["kfbleh"],
          enterpriseAppId: "mklfypyujwumgwdzae",
          provisioningState: "Succeeded",
          singleSignOnState: "Initial",
          singleSignOnUrl: "ymmtzkyghvinvhgnqlzwrr",
        },
        workspaceId: "xsepuskdhejaadusyxq",
        workspaceName: "L.-y_--:",
      },
      user: {
        emailAddress: ".K_@e7N-g1.xjqnbPs",
        firstName: "qeuofehzypzljgcuysugefbgxde",
        lastName: "g",
        phoneNumber: "aqpyxznvqpgkzohevynofrjdfgoo",
        upn: "uwtprzdfpsqmktx",
      },
    },
    tags: { key1474: "bqqyipxnbbxryhznyaosmtpo" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AstroManagementClient(credential, subscriptionId);
  const result = await client.organizations.beginUpdateAndWait(
    resourceGroupName,
    organizationName,
    properties,
  );
  console.log(result);
}
