using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.BotService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.BotService;

// Generated from example definition: specification/botservice/resource-manager/Microsoft.BotService/stable/2022-09-15/examples/CheckNameAvailability.json
// this example is just showing the usage of "Bots_GetCheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
BotServiceNameAvailabilityContent content = new BotServiceNameAvailabilityContent
{
    Name = "testbotname",
    ResourceType = new ResourceType("string"),
};
BotServiceNameAvailabilityResult result = await tenantResource.CheckBotServiceNameAvailabilityAsync(content);

Console.WriteLine($"Succeeded: {result}");
