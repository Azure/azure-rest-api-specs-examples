using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeploymentManager;
using Azure.ResourceManager.DeploymentManager.Models;

// Generated from example definition: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/step_health_check_createorupdate.json
// this example is just showing the usage of "Steps_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StepResource created on azure
// for more information of creating StepResource, please refer to the document of StepResource
string subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
string resourceGroupName = "myResourceGroup";
string stepName = "healthCheckStep";
ResourceIdentifier stepResourceId = StepResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, stepName);
StepResource stepResource = client.GetStepResource(stepResourceId);

// invoke the operation
StepResourceData data = new StepResourceData(new AzureLocation("centralus"), new HealthCheckStepProperties(new RestHealthCheckStepAttributes(XmlConvert.ToTimeSpan("PT2H"))
{
    HealthChecks =
    {
    new RestHealthCheck("appHealth",new RestRequest(RestRequestMethod.GET,new Uri("https://resthealth.healthservice.com/api/applications/contosoApp/healthStatus"),new ApiKeyAuthentication("Code",RestAuthLocation.Query,"NBCapiMOBQyAAbCkeytoPadnvO0eGHmidwFz5rXpappznKp3Jt7LLg==")))
    {
    Response = new RestResponse()
    {
    SuccessStatusCodes =
    {
    "OK"
    },
    Regex = new RestResponseRegex()
    {
    Matches =
    {
    "(?i)Contoso-App","(?i)\"health_status\":((.|\n)*)\"(green|yellow)\"","(?mi)^(\"application_host\": 94781052)$"
    },
    MatchQuantifier = RestMatchQuantifier.All,
    },
    },
    },new RestHealthCheck("serviceHealth",new RestRequest(RestRequestMethod.GET,new Uri("https://resthealth.healthservice.com/api/services/contosoService/healthStatus"),new ApiKeyAuthentication("code",RestAuthLocation.Header,"NBCapiMOBQyAAbCkeytoPadnvO0eGHmidwFz5rXpappznKp3Jt7LLg==")))
    {
    Response = new RestResponse()
    {
    SuccessStatusCodes =
    {
    "OK"
    },
    Regex = new RestResponseRegex()
    {
    Matches =
    {
    "(?i)Contoso-Service-EndToEnd","(?i)\"health_status\":((.|\n)*)\"(green)\""
    },
    MatchQuantifier = RestMatchQuantifier.All,
    },
    },
    }
    },
    WaitDuration = XmlConvert.ToTimeSpan("PT15M"),
    MaxElasticDuration = XmlConvert.ToTimeSpan("PT30M"),
}))
{
    Tags =
    {
    },
};
ArmOperation<StepResource> lro = await stepResource.UpdateAsync(WaitUntil.Completed, data);
StepResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
StepResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
