using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DataMigration.Models;
using Azure.ResourceManager.DataMigration;

// Generated from example definition: specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2022-03-30-preview/examples/Tasks_CreateOrUpdate.json
// this example is just showing the usage of "Tasks_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ProjectResource created on azure
// for more information of creating ProjectResource, please refer to the document of ProjectResource
string subscriptionId = "fc04246f-04c5-437e-ac5e-206a19e7193f";
string groupName = "DmsSdkRg";
string serviceName = "DmsSdkService";
string projectName = "DmsSdkProject";
ResourceIdentifier projectResourceId = ProjectResource.CreateResourceIdentifier(subscriptionId, groupName, serviceName, projectName);
ProjectResource project = client.GetProjectResource(projectResourceId);

// get the collection of this ServiceProjectTaskResource
ServiceProjectTaskCollection collection = project.GetServiceProjectTasks();

// invoke the operation
string taskName = "DmsSdkTask";
ProjectTaskData data = new ProjectTaskData
{
    Properties = new ConnectToTargetSqlDBTaskProperties
    {
        Input = new ConnectToTargetSqlDBTaskInput(new SqlConnectionInfo("ssma-test-server.database.windows.net")
        {
            Authentication = AuthenticationType.SqlAuthentication,
            EncryptConnection = true,
            TrustServerCertificate = true,
            UserName = "testuser",
            Password = "testpassword",
        }),
    },
};
ArmOperation<ServiceProjectTaskResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, taskName, data);
ServiceProjectTaskResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ProjectTaskData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
