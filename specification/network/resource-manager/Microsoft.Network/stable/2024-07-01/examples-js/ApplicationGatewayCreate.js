const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates the specified application gateway.
 *
 * @summary Creates or updates the specified application gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ApplicationGatewayCreate.json
 */
async function createApplicationGateway() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const applicationGatewayName = "appgw";
  const parameters = {
    backendAddressPools: [
      {
        name: "appgwpool",
        backendAddresses: [{ ipAddress: "10.0.1.1" }, { ipAddress: "10.0.1.2" }],
      },
      {
        name: "appgwpool1",
        backendAddresses: [{ ipAddress: "10.0.0.1" }, { ipAddress: "10.0.0.2" }],
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool1",
      },
    ],
    backendHttpSettingsCollection: [
      {
        name: "appgwbhs",
        cookieBasedAffinity: "Disabled",
        port: 80,
        requestTimeout: 30,
        protocol: "Http",
      },
    ],
    frontendIPConfigurations: [
      {
        name: "appgwfip",
        publicIPAddress: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/appgwpip",
        },
      },
    ],
    frontendPorts: [
      { name: "appgwfp", port: 443 },
      { name: "appgwfp80", port: 80 },
    ],
    gatewayIPConfigurations: [
      {
        name: "appgwipc",
        subnet: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet/subnets/appgwsubnet",
        },
      },
    ],
    globalConfiguration: {
      enableRequestBuffering: true,
      enableResponseBuffering: true,
    },
    httpListeners: [
      {
        name: "appgwhl",
        frontendIPConfiguration: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip",
        },
        frontendPort: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp",
        },
        requireServerNameIndication: false,
        sslCertificate: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslCertificates/sslcert",
        },
        sslProfile: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/sslProfiles/sslProfile1",
        },
        protocol: "Https",
      },
      {
        name: "appgwhttplistener",
        frontendIPConfiguration: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendIPConfigurations/appgwfip",
        },
        frontendPort: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/frontendPorts/appgwfp80",
        },
        protocol: "Http",
      },
    ],
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/subid/resourceGroups/rg1/providers/MicrosoftManagedIdentity/userAssignedIdentities/identity1":
          {},
      },
    },
    location: "eastus",
    requestRoutingRules: [
      {
        name: "appgwrule",
        backendAddressPool: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendAddressPools/appgwpool",
        },
        backendHttpSettings: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/backendHttpSettingsCollection/appgwbhs",
        },
        httpListener: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/httpListeners/appgwhl",
        },
        priority: 10,
        rewriteRuleSet: {
          id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/rewriteRuleSets/rewriteRuleSet1",
        },
        ruleType: "Basic",
      },
    ],
    rewriteRuleSets: [
      {
        name: "rewriteRuleSet1",
        rewriteRules: [
          {
            name: "Set X-Forwarded-For",
            actionSet: {
              requestHeaderConfigurations: [
                {
                  headerName: "X-Forwarded-For",
                  headerValue: "{var_add_x_forwarded_for_proxy}",
                },
              ],
              responseHeaderConfigurations: [
                {
                  headerName: "Strict-Transport-Security",
                  headerValue: "max-age=31536000",
                },
              ],
              urlConfiguration: { modifiedPath: "/abc" },
            },
            conditions: [
              {
                ignoreCase: true,
                negate: false,
                pattern: "^Bearer",
                variable: "http_req_Authorization",
              },
            ],
            ruleSequence: 102,
          },
        ],
      },
    ],
    sku: { name: "Standard_v2", capacity: 3, tier: "Standard_v2" },
    sslCertificates: [
      { name: "sslcert", data: "****", password: "****" },
      { name: "sslcert2", keyVaultSecretId: "https://kv/secret" },
    ],
    sslProfiles: [
      {
        name: "sslProfile1",
        clientAuthConfiguration: { verifyClientCertIssuerDN: true },
        sslPolicy: {
          cipherSuites: ["TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256"],
          minProtocolVersion: "TLSv1_1",
          policyType: "Custom",
        },
        trustedClientCertificates: [
          {
            id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/applicationGateways/appgw/trustedClientCertificates/clientcert",
          },
        ],
      },
    ],
    trustedClientCertificates: [{ name: "clientcert", data: "****" }],
    trustedRootCertificates: [
      { name: "rootcert", data: "****" },
      { name: "rootcert1", keyVaultSecretId: "https://kv/secret" },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.applicationGateways.beginCreateOrUpdateAndWait(
    resourceGroupName,
    applicationGatewayName,
    parameters,
  );
  console.log(result);
}
