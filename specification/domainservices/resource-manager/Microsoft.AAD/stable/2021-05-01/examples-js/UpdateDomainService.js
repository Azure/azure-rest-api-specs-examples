const { DomainServicesResourceProvider } = require("@azure/arm-domainservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Update Domain Service operation can be used to update the existing deployment. The update call only supports the properties listed in the PATCH body.
 *
 * @summary The Update Domain Service operation can be used to update the existing deployment. The update call only supports the properties listed in the PATCH body.
 * x-ms-original-file: specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/UpdateDomainService.json
 */
async function updateDomainService() {
  const subscriptionId = "1639790a-76a2-4ac4-98d9-8562f5dfcb4d";
  const resourceGroupName = "TestResourceGroup";
  const domainServiceName = "TestDomainService.com";
  const domainService = {
    configDiagnostics: {
      lastExecuted: new Date("2021-05-05T12:00:23Z;"),
      validatorResults: [
        {
          issues: [{ descriptionParams: [], id: "AADDS-CFG-DIAG-I20" }],
          replicaSetSubnetDisplayName: "West US/aadds-subnet",
          status: "Warning",
          validatorId: "AADDS-CFG-DIAG-V06",
        },
      ],
    },
    domainSecuritySettings: {
      ntlmV1: "Enabled",
      syncNtlmPasswords: "Enabled",
      tlsV1: "Disabled",
    },
    filteredSync: "Enabled",
    ldapsSettings: {
      externalAccess: "Enabled",
      ldaps: "Enabled",
      pfxCertificate: "MIIDPDCCAiSgAwIBAgIQQUI9P6tq2p9OFIJa7DLNvTANBgkqhkiG9w0BAQsFADAgMR4w...",
      pfxCertificatePassword: "<pfxCertificatePassword>",
    },
    notificationSettings: {
      additionalRecipients: ["jicha@microsoft.com", "caalmont@microsoft.com"],
      notifyDcAdmins: "Enabled",
      notifyGlobalAdmins: "Enabled",
    },
    replicaSets: [
      {
        location: "West US",
        subnetId:
          "/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/TestNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/TestVnetWUS/subnets/TestSubnetWUS",
      },
      {
        location: "East US",
        subnetId:
          "/subscriptions/1639790a-76a2-4ac4-98d9-8562f5dfcb4d/resourceGroups/TestNetworkResourceGroup/providers/Microsoft.Network/virtualNetworks/TestVnetEUS/subnets/TestSubnetEUS",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new DomainServicesResourceProvider(credential, subscriptionId);
  const result = await client.domainServices.beginUpdateAndWait(
    resourceGroupName,
    domainServiceName,
    domainService
  );
  console.log(result);
}

updateDomainService().catch(console.error);
