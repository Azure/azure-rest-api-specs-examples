using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.BotService;
using Azure.ResourceManager.BotService.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/preview/2022-06-15-preview/examples/CheckNameAvailability.json
// this example is just showing the usage of "Bots_GetCheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
CheckNameAvailabilityRequestBody checkNameAvailabilityRequestBody = new CheckNameAvailabilityRequestBody()
{
    Name = "testbotname",
    ResourceType = "string",
};
CheckNameAvailabilityResponseBody result = await tenantResource.GetCheckNameAvailabilityBotAsync(checkNameAvailabilityRequestBody);

Console.WriteLine($"Succeeded: {result}");
