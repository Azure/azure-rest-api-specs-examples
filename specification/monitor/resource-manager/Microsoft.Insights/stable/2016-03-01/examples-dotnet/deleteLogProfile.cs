using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Monitor;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/deleteLogProfile.json
// this example is just showing the usage of "LogProfiles_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogProfileResource created on azure
// for more information of creating LogProfileResource, please refer to the document of LogProfileResource
string subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
string logProfileName = "Rac46PostSwapRG";
ResourceIdentifier logProfileResourceId = LogProfileResource.CreateResourceIdentifier(subscriptionId, logProfileName);
LogProfileResource logProfile = client.GetLogProfileResource(logProfileResourceId);

// invoke the operation
await logProfile.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
