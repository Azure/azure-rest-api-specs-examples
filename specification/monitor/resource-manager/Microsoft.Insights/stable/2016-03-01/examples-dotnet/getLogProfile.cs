using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Monitor;
using Azure.ResourceManager.Monitor.Models;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/getLogProfile.json
// this example is just showing the usage of "LogProfiles_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogProfileResource created on azure
// for more information of creating LogProfileResource, please refer to the document of LogProfileResource
string subscriptionId = "df602c9c-7aa0-407d-a6fb-eb20c8bd1192";
string logProfileName = "default";
ResourceIdentifier logProfileResourceId = LogProfileResource.CreateResourceIdentifier(subscriptionId, logProfileName);
LogProfileResource logProfile = client.GetLogProfileResource(logProfileResourceId);

// invoke the operation
LogProfileResource result = await logProfile.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LogProfileData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
