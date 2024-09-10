using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.HDInsight.Models;
using Azure.ResourceManager.HDInsight;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/GetApplicationInProgress.json
// this example is just showing the usage of "Applications_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HDInsightApplicationResource created on azure
// for more information of creating HDInsightApplicationResource, please refer to the document of HDInsightApplicationResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string clusterName = "cluster1";
string applicationName = "app";
ResourceIdentifier hdInsightApplicationResourceId = HDInsightApplicationResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, clusterName, applicationName);
HDInsightApplicationResource hdInsightApplication = client.GetHDInsightApplicationResource(hdInsightApplicationResourceId);

// invoke the operation
HDInsightApplicationResource result = await hdInsightApplication.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HDInsightApplicationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
