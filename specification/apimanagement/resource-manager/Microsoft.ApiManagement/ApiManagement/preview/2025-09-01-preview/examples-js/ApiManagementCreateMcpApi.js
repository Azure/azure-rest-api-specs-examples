const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates new or updates existing specified API of the API Management service instance.
 *
 * @summary creates new or updates existing specified API of the API Management service instance.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateMcpApi.json
 */
async function apiManagementCreateMcpApi() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.api.createOrUpdate("rg1", "apimService1", "mcp-api", {
    apiType: "mcp",
    path: "mcp-api",
    description: "MCP API for AI agent tool discovery and invocation",
    displayName: "MCP API",
    protocols: ["https"],
    serviceUrl: "https://mcp-backend.contoso.com",
    mcpProperties: {
      transportType: "streamable",
      endpoints: [{ name: "message", uriTemplate: "/mcp/messages" }],
    },
  });
  console.log(result);
}
