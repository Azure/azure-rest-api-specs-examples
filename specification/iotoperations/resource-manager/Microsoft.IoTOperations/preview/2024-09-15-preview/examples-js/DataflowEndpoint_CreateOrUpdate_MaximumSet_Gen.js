const { IoTOperationsClient } = require("@azure/arm-iotoperations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a DataflowEndpointResource
 *
 * @summary create a DataflowEndpointResource
 * x-ms-original-file: 2024-09-15-preview/DataflowEndpoint_CreateOrUpdate_MaximumSet_Gen.json
 */
async function dataflowEndpointCreateOrUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "F8C729F9-DF9C-4743-848F-96EE433D8E53";
  const client = new IoTOperationsClient(credential, subscriptionId);
  const result = await client.dataflowEndpoint.createOrUpdate(
    "rgiotoperations",
    "resource-name123",
    "resource-name123",
    {
      properties: {
        endpointType: "DataExplorer",
        dataExplorerSettings: {
          authentication: {
            method: "SystemAssignedManagedIdentity",
            systemAssignedManagedIdentitySettings: {
              audience: "psxomrfbhoflycm",
            },
            userAssignedManagedIdentitySettings: {
              clientId: "fb90f267-8872-431a-a76a-a1cec5d3c4d2",
              scope: "zop",
              tenantId: "ed060aa2-71ff-4d3f-99c4-a9138356fdec",
            },
          },
          database: "yqcdpjsifm",
          host: "<cluster>.<region>.kusto.windows.net",
          batching: { latencySeconds: 9312, maxMessages: 9028 },
        },
        dataLakeStorageSettings: {
          authentication: {
            method: "SystemAssignedManagedIdentity",
            accessTokenSettings: { secretRef: "sevriyphcvnlrnfudqzejecwa" },
            systemAssignedManagedIdentitySettings: {
              audience: "psxomrfbhoflycm",
            },
            userAssignedManagedIdentitySettings: {
              clientId: "fb90f267-8872-431a-a76a-a1cec5d3c4d2",
              scope: "zop",
              tenantId: "ed060aa2-71ff-4d3f-99c4-a9138356fdec",
            },
          },
          host: "<account>.blob.core.windows.net",
          batching: { latencySeconds: 9312, maxMessages: 9028 },
        },
        fabricOneLakeSettings: {
          authentication: {
            method: "SystemAssignedManagedIdentity",
            systemAssignedManagedIdentitySettings: {
              audience: "psxomrfbhoflycm",
            },
            userAssignedManagedIdentitySettings: {
              clientId: "fb90f267-8872-431a-a76a-a1cec5d3c4d2",
              scope: "zop",
              tenantId: "ed060aa2-71ff-4d3f-99c4-a9138356fdec",
            },
          },
          names: { lakehouseName: "wpeathi", workspaceName: "nwgmitkbljztgms" },
          oneLakePathType: "Files",
          host: "https://<host>.fabric.microsoft.com",
          batching: { latencySeconds: 9312, maxMessages: 9028 },
        },
        kafkaSettings: {
          authentication: {
            method: "SystemAssignedManagedIdentity",
            systemAssignedManagedIdentitySettings: {
              audience: "psxomrfbhoflycm",
            },
            userAssignedManagedIdentitySettings: {
              clientId: "fb90f267-8872-431a-a76a-a1cec5d3c4d2",
              scope: "zop",
              tenantId: "ed060aa2-71ff-4d3f-99c4-a9138356fdec",
            },
            saslSettings: {
              saslType: "Plain",
              secretRef: "visyxoztqnylvbyokhtmpdkwes",
            },
            x509CertificateSettings: { secretRef: "afwizrystfslkfqd" },
          },
          consumerGroupId: "ukkzcjiyenhxokat",
          host: "pwcqfiqclcgneolpewnyavoulbip",
          batching: {
            mode: "Enabled",
            latencyMs: 3679,
            maxBytes: 8887,
            maxMessages: 2174,
          },
          copyMqttProperties: "Enabled",
          compression: "None",
          kafkaAcks: "Zero",
          partitionStrategy: "Default",
          tls: {
            mode: "Enabled",
            trustedCaCertificateConfigMapRef: "tectjjvukvelsreihwadh",
          },
        },
        localStorageSettings: { persistentVolumeClaimRef: "jjwqwvd" },
        mqttSettings: {
          authentication: {
            method: "SystemAssignedManagedIdentity",
            systemAssignedManagedIdentitySettings: {
              audience: "psxomrfbhoflycm",
            },
            userAssignedManagedIdentitySettings: {
              clientId: "fb90f267-8872-431a-a76a-a1cec5d3c4d2",
              scope: "zop",
              tenantId: "ed060aa2-71ff-4d3f-99c4-a9138356fdec",
            },
            serviceAccountTokenSettings: { audience: "ejbklrbxgjaqleoycgpje" },
            x509CertificateSettings: { secretRef: "afwizrystfslkfqd" },
          },
          clientIdPrefix: "kkljsdxdirfhwxtkavldekeqhv",
          host: "nyhnxqnbspstctl",
          protocol: "Mqtt",
          keepAliveSeconds: 0,
          retain: "Keep",
          maxInflightMessages: 0,
          qos: 1,
          sessionExpirySeconds: 0,
          tls: {
            mode: "Enabled",
            trustedCaCertificateConfigMapRef: "tectjjvukvelsreihwadh",
          },
        },
      },
      extendedLocation: {
        name: "qmbrfwcpwwhggszhrdjv",
        type: "CustomLocation",
      },
    },
  );
  console.log(result);
}
