using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automanage;

// Generated from example definition: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfileVersions.json
// this example is just showing the usage of "ConfigurationProfilesVersions_ListChildResources" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomanageConfigurationProfileResource created on azure
// for more information of creating AutomanageConfigurationProfileResource, please refer to the document of AutomanageConfigurationProfileResource
string subscriptionId = "mySubscriptionId";
string resourceGroupName = "myResourceGroupName";
string configurationProfileName = "customConfigurationProfile";
ResourceIdentifier automanageConfigurationProfileResourceId = AutomanageConfigurationProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, configurationProfileName);
AutomanageConfigurationProfileResource automanageConfigurationProfile = client.GetAutomanageConfigurationProfileResource(automanageConfigurationProfileResourceId);

// get the collection of this AutomanageConfigurationProfileVersionResource
AutomanageConfigurationProfileVersionCollection collection = automanageConfigurationProfile.GetAutomanageConfigurationProfileVersions();

// invoke the operation and iterate over the result
await foreach (AutomanageConfigurationProfileVersionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AutomanageConfigurationProfileData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
