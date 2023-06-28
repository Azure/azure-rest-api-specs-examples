using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HDInsight;
using Azure.ResourceManager.HDInsight.Models;

// Generated from example definition: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/GetApplicationCreationAsyncOperationStatus.json
// this example is just showing the usage of "Applications_GetAzureAsyncOperationStatus" operation, for the dependent resources, they will have to be created separately.

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
string operationId = "CF938302-6B4D-44A0-A6D2-C0D67E847AEC";
HDInsightAsyncOperationResult result = await hdInsightApplication.GetAzureAsyncOperationStatusAsync(operationId);

Console.WriteLine($"Succeeded: {result}");
