using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Automanage;

// Generated from example definition: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfileVersion.json
// this example is just showing the usage of "ConfigurationProfilesVersions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AutomanageConfigurationProfileVersionResource created on azure
// for more information of creating AutomanageConfigurationProfileVersionResource, please refer to the document of AutomanageConfigurationProfileVersionResource
string subscriptionId = "subid";
string resourceGroupName = "rg";
string configurationProfileName = "customConfigurationProfile";
string versionName = "version1";
ResourceIdentifier automanageConfigurationProfileVersionResourceId = AutomanageConfigurationProfileVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, configurationProfileName, versionName);
AutomanageConfigurationProfileVersionResource automanageConfigurationProfileVersion = client.GetAutomanageConfigurationProfileVersionResource(automanageConfigurationProfileVersionResourceId);

// invoke the operation
await automanageConfigurationProfileVersion.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
