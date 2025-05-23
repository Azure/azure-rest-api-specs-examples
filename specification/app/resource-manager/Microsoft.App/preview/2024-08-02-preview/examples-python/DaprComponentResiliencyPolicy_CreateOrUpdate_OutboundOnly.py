from azure.identity import DefaultAzureCredential

from azure.mgmt.appcontainers import ContainerAppsAPIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-appcontainers
# USAGE
    python dapr_component_resiliency_policy_create_or_update_outbound_only.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ContainerAppsAPIClient(
        credential=DefaultAzureCredential(),
        subscription_id="8efdecc5-919e-44eb-b179-915dca89ebf9",
    )

    response = client.dapr_component_resiliency_policies.create_or_update(
        resource_group_name="examplerg",
        environment_name="myenvironment",
        component_name="mydaprcomponent",
        name="myresiliencypolicy",
        dapr_component_resiliency_policy_envelope={
            "properties": {
                "outboundPolicy": {
                    "circuitBreakerPolicy": {"consecutiveErrors": 3, "intervalInSeconds": 60, "timeoutInSeconds": 20},
                    "httpRetryPolicy": {
                        "maxRetries": 5,
                        "retryBackOff": {"initialDelayInMilliseconds": 100, "maxIntervalInMilliseconds": 30000},
                    },
                    "timeoutPolicy": {"responseTimeoutInSeconds": 12},
                }
            }
        },
    )
    print(response)


# x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/DaprComponentResiliencyPolicy_CreateOrUpdate_OutboundOnly.json
if __name__ == "__main__":
    main()
