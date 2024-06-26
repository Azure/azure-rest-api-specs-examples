from azure.identity import DefaultAzureCredential
from azure.mgmt.machinelearningservices import MachineLearningServicesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-machinelearningservices
# USAGE
    python compute_instance.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = MachineLearningServicesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
    )

    response = client.compute.begin_create_or_update(
        resource_group_name="testrg123",
        workspace_name="workspaces123",
        compute_name="compute123",
        parameters={
            "location": "eastus",
            "properties": {
                "computeType": "ComputeInstance",
                "properties": {
                    "applicationSharingPolicy": "Personal",
                    "computeInstanceAuthorizationType": "personal",
                    "customServices": [
                        {
                            "docker": {"privileged": True},
                            "endpoints": [{"name": "connect", "protocol": "http", "published": 8787, "target": 8787}],
                            "environmentVariables": {"test_variable": {"type": "local", "value": "test_value"}},
                            "image": {"reference": "ghcr.io/azure/rocker-rstudio-ml-verse:latest", "type": "docker"},
                            "name": "rstudio",
                            "volumes": [
                                {
                                    "readOnly": False,
                                    "source": "/home/azureuser/cloudfiles",
                                    "target": "/home/azureuser/cloudfiles",
                                    "type": "bind",
                                }
                            ],
                        }
                    ],
                    "personalComputeInstanceSettings": {
                        "assignedUser": {
                            "objectId": "00000000-0000-0000-0000-000000000000",
                            "tenantId": "00000000-0000-0000-0000-000000000000",
                        }
                    },
                    "sshSettings": {"sshPublicAccess": "Disabled"},
                    "subnet": {"id": "test-subnet-resource-id"},
                    "vmSize": "STANDARD_NC6",
                },
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2023-04-01/examples/Compute/createOrUpdate/ComputeInstance.json
if __name__ == "__main__":
    main()
