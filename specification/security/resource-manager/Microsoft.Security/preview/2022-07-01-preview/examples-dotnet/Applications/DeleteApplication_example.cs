using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/DeleteApplication_example.json
// this example is just showing the usage of "Application_Delete" operation, for the dependent resources, they will have to be created separately.

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
await subscriptionSecurityApplication.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
