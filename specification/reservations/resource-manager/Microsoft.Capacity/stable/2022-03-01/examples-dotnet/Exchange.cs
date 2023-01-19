using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Reservations.Models;

// Generated from example definition: specification/reservations/resource-manager/Microsoft.Capacity/stable/2022-03-01/examples/Exchange.json
// this example is just showing the usage of "Exchange_Post" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
ExchangeContent content = new ExchangeContent()
{
    ExchangeRequestSessionId = Guid.Parse("66e2ac8f-439e-4345-8235-6fef07608081"),
};
ArmOperation<ExchangeResult> lro = await tenantResource.ExchangeAsync(WaitUntil.Completed, content);
ExchangeResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
