using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Automanage.Models;
using Azure.ResourceManager.Automanage;

// Generated from example definition: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfile.json
// this example is just showing the usage of "ConfigurationProfiles_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomanageConfigurationProfileResource created on azure
// for more information of creating AutomanageConfigurationProfileResource, please refer to the document of AutomanageConfigurationProfileResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string configurationProfileName = "customConfigurationProfile";
ResourceIdentifier automanageConfigurationProfileResourceId = AutomanageConfigurationProfileResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, configurationProfileName);
AutomanageConfigurationProfileResource automanageConfigurationProfile = client.GetAutomanageConfigurationProfileResource(automanageConfigurationProfileResourceId);

// invoke the operation
await automanageConfigurationProfile.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
