using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.RecoveryServices.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.RecoveryServices;

// Generated from example definition: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/CheckNameAvailability_Available.json
// this example is just showing the usage of "RecoveryServices_CheckNameAvailability" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "77777777-b0c6-47a2-b37c-d8e65a629c18";
string resourceGroupName = "resGroupFoo";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("westus");
RecoveryServicesNameAvailabilityContent content = new RecoveryServicesNameAvailabilityContent()
{
    ResourceType = new ResourceType("Microsoft.RecoveryServices/Vaults"),
    Name = "swaggerExample",
};
RecoveryServicesNameAvailabilityResult result = await resourceGroupResource.CheckRecoveryServicesNameAvailabilityAsync(location, content);

Console.WriteLine($"Succeeded: {result}");
