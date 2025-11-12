const { DeviceRegistryManagementClient } = require("@azure/arm-deviceregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a NamespaceDevice
 *
 * @summary create a NamespaceDevice
 * x-ms-original-file: 2025-10-01/CreateOrReplace_NamespaceDevice_Edge_x509.json
 */
async function createEdgeEnabledDeviceWithX509InboundAuthentication() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new DeviceRegistryManagementClient(credential, subscriptionId);
  const result = await client.namespaceDevices.createOrReplace(
    "myResourceGroup",
    "adr-namespace-gbk0925-n01",
    "namespace-device-on-edge",
    {
      extendedLocation: {
        type: "CustomLocation",
        name: "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/myResourceGroup/providers/microsoft.extendedlocation/customlocations/location1",
      },
      location: "West Europe",
      properties: {
        endpoints: {
          inbound: {
            theV1OPCUAEndpoint: {
              address: "opc.tcp://192.168.86.23:51211/UA/SampleServer",
              endpointType: "microsoft.opcua",
              version: "2",
              authentication: {
                method: "Certificate",
                x509Credentials: {
                  certificateSecretName: "cert-secret",
                  keySecretName: "key-secret",
                  intermediateCertificatesSecretName: "intermediate-certs-secret",
                },
              },
            },
            theV2OPCUAEndpoint: {
              address: "opc.tcp://192.168.86.23:51211/UA/SampleServer",
              endpointType: "microsoft.opcua",
              version: "2",
              authentication: {
                method: "Certificate",
                x509Credentials: { certificateSecretName: "cert-secret" },
              },
              trustSettings: { trustList: "trust-secret-reference" },
            },
          },
        },
        externalDeviceId: "unique-edge-device-identifier",
        enabled: true,
        attributes: {
          deviceType: "OPCUAServers",
          deviceOwner: "OT",
          deviceCategory: 16,
        },
      },
    },
  );
  console.log(result);
}
