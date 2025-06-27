using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CloudHealth.Models;
using Azure.ResourceManager.CloudHealth;

// Generated from example definition: 2025-05-01-preview/Relationships_Delete.json
// this example is just showing the usage of "Relationship_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HealthModelRelationshipResource created on azure
// for more information of creating HealthModelRelationshipResource, please refer to the document of HealthModelRelationshipResource
string subscriptionId = "4980D7D5-4E07-47AD-AD34-E76C6BC9F061";
string resourceGroupName = "rgopenapi";
string healthModelName = "model1";
string relationshipName = "rel1";
ResourceIdentifier healthModelRelationshipResourceId = HealthModelRelationshipResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, healthModelName, relationshipName);
HealthModelRelationshipResource healthModelRelationship = client.GetHealthModelRelationshipResource(healthModelRelationshipResourceId);

// invoke the operation
await healthModelRelationship.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
