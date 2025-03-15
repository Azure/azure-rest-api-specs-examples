using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataShare.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.DataShare;

// Generated from example definition: specification/datashare/resource-manager/Microsoft.DataShare/stable/2021-08-01/examples/EmailRegistrations_ActivateEmail.json
// this example is just showing the usage of "EmailRegistrations_ActivateEmail" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
AzureLocation location = new AzureLocation("East US 2");
DataShareEmailRegistration emailRegistration = new DataShareEmailRegistration
{
    ActivationCode = "djsfhakj2lekowd3wepfklpwe9lpflcd",
};
DataShareEmailRegistration result = await tenantResource.ActivateEmailAsync(location, emailRegistration);

Console.WriteLine($"Succeeded: {result}");
