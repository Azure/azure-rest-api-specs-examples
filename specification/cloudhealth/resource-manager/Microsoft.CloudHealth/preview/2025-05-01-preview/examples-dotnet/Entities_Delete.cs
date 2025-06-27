using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CloudHealth.Models;
using Azure.ResourceManager.CloudHealth;

// Generated from example definition: 2025-05-01-preview/Entities_Delete.json
// this example is just showing the usage of "Entity_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthModelEntityResource created on azure
// for more information of creating HealthModelEntityResource, please refer to the document of HealthModelEntityResource
string subscriptionId = "4980D7D5-4E07-47AD-AD34-E76C6BC9F061";
string resourceGroupName = "rgopenapi";
string healthModelName = "model1";
string entityName = "U4VTRFlUkm9kR6H23-c-6U-XHq7n";
ResourceIdentifier healthModelEntityResourceId = HealthModelEntityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, healthModelName, entityName);
HealthModelEntityResource healthModelEntity = client.GetHealthModelEntityResource(healthModelEntityResourceId);

// invoke the operation
await healthModelEntity.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
