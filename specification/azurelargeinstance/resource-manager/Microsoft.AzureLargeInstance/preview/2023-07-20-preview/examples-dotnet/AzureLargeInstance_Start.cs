using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.LargeInstance;
using Azure.ResourceManager.LargeInstance.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/azurelargeinstance/resource-manager/Microsoft.AzureLargeInstance/preview/2023-07-20-preview/examples/AzureLargeInstance_Start.json
// this example is just showing the usage of "AzureLargeInstance_Start" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LargeInstanceResource created on azure
// for more information of creating LargeInstanceResource, please refer to the document of LargeInstanceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string azureLargeInstanceName = "myALInstance";
ResourceIdentifier largeInstanceResourceId = LargeInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, azureLargeInstanceName);
LargeInstanceResource largeInstance = client.GetLargeInstanceResource(largeInstanceResourceId);

// invoke the operation
ArmOperation<LargeInstanceOperationStatusResult> lro = await largeInstance.StartAsync(WaitUntil.Completed);
LargeInstanceOperationStatusResult result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
