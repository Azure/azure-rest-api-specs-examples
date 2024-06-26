from azure.identity import DefaultAzureCredential
from azure.mgmt.machinelearningservices import MachineLearningServicesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-machinelearningservices
# USAGE
    python create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MachineLearningServicesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-1111-2222-3333-444444444444",
    )

    response = client.online_deployments.begin_create_or_update(
        resource_group_name="test-rg",
        workspace_name="my-aml-workspace",
        endpoint_name="testEndpointName",
        deployment_name="testDeploymentName",
        body={
            "identity": {"type": "SystemAssigned", "userAssignedIdentities": {"string": {}}},
            "kind": "string",
            "location": "string",
            "properties": {
                "appInsightsEnabled": False,
                "codeConfiguration": {"codeId": "string", "scoringScript": "string"},
                "description": "string",
                "endpointComputeType": "Managed",
                "environmentId": "string",
                "environmentVariables": {"string": "string"},
                "instanceType": "string",
                "livenessProbe": {
                    "failureThreshold": 1,
                    "initialDelay": "PT5M",
                    "period": "PT5M",
                    "successThreshold": 1,
                    "timeout": "PT5M",
                },
                "model": "string",
                "modelMountPath": "string",
                "properties": {"string": "string"},
                "readinessProbe": {
                    "failureThreshold": 30,
                    "initialDelay": "PT1S",
                    "period": "PT10S",
                    "successThreshold": 1,
                    "timeout": "PT2S",
                },
                "requestSettings": {
                    "maxConcurrentRequestsPerInstance": 1,
                    "maxQueueWait": "PT5M",
                    "requestTimeout": "PT5M",
                },
                "scaleSettings": {"scaleType": "Default"},
            },
            "sku": {"capacity": 1, "family": "string", "name": "string", "size": "string", "tier": "Free"},
            "tags": {},
        },
    ).result()
    print(response)


# x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2023-04-01/examples/OnlineDeployment/ManagedOnlineDeployment/createOrUpdate.json
if __name__ == "__main__":
    main()
