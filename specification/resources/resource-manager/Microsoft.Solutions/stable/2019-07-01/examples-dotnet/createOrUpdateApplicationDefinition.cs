using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/resources/resource-manager/Microsoft.Solutions/stable/2019-07-01/examples/createOrUpdateApplicationDefinition.json
// this example is just showing the usage of "ApplicationDefinitions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ArmApplicationDefinitionResource created on azure
// for more information of creating ArmApplicationDefinitionResource, please refer to the document of ArmApplicationDefinitionResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string applicationDefinitionName = "myManagedApplicationDef";
ResourceIdentifier armApplicationDefinitionResourceId = ArmApplicationDefinitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, applicationDefinitionName);
ArmApplicationDefinitionResource armApplicationDefinition = client.GetArmApplicationDefinitionResource(armApplicationDefinitionResourceId);

// invoke the operation
ArmApplicationDefinitionData data = new ArmApplicationDefinitionData(new AzureLocation("East US 2"), ArmApplicationLockLevel.None)
{
    DisplayName = "myManagedApplicationDef",
    Authorizations = { new ArmApplicationAuthorization(Guid.Parse("validprincipalguid"), "validroleguid") },
    Description = "myManagedApplicationDef description",
    PackageFileUri = new Uri("https://path/to/packagezipfile"),
};
ArmOperation<ArmApplicationDefinitionResource> lro = await armApplicationDefinition.UpdateAsync(WaitUntil.Completed, data);
ArmApplicationDefinitionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ArmApplicationDefinitionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
