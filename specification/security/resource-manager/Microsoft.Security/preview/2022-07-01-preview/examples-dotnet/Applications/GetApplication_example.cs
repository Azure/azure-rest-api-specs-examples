using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/GetApplication_example.json
// this example is just showing the usage of "Application_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionSecurityApplicationResource created on azure
// for more information of creating SubscriptionSecurityApplicationResource, please refer to the document of SubscriptionSecurityApplicationResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
string applicationId = "ad9a8e26-29d9-4829-bb30-e597a58cdbb8";
ResourceIdentifier subscriptionSecurityApplicationResourceId = SubscriptionSecurityApplicationResource.CreateResourceIdentifier(subscriptionId, applicationId);
SubscriptionSecurityApplicationResource subscriptionSecurityApplication = client.GetSubscriptionSecurityApplicationResource(subscriptionSecurityApplicationResourceId);

// invoke the operation
SubscriptionSecurityApplicationResource result = await subscriptionSecurityApplication.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityApplicationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
