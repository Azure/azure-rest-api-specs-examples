from azure.identity import DefaultAzureCredential

from azure.mgmt.redhatopenshift import AzureRedHatOpenShiftClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-redhatopenshift
# USAGE
    python machine_pools_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureRedHatOpenShiftClient(
        credential=DefaultAzureCredential(),
        subscription_id="subscriptionId",
    )

    response = client.machine_pools.create_or_update(
        resource_group_name="resourceGroup",
        resource_name="resourceName",
        child_resource_name="childResourceName",
        parameters={
            "properties": {
                "resources": "ewogICAgImFwaVZlcnNpb24iOiAiaGl2ZS5vcGVuc2hpZnQuaW8vdjEiLAogICAgImtpbmQiOiAiTWFjaGluZVBvb2wiLAogICAgIm1ldGFkYXRhIjogewogICAgICAgICJuYW1lIjogInRlc3QtY2x1c3Rlci13b3JrZXIiLAogICAgICAgICJuYW1lc3BhY2UiOiAiYXJvLWY2MGFlOGEyLWJjYTEtNDk4Ny05MDU2LVhYWFhYWFhYWFhYWCIKICAgIH0sCiAgICAic3BlYyI6IHsKICAgICAgICAiY2x1c3RlckRlcGxveW1lbnRSZWYiOiB7CiAgICAgICAgICAgICJuYW1lIjogInRlc3QtY2x1c3RlciIKICAgICAgICB9LAogICAgICAgICJuYW1lIjogIndvcmtlciIsCiAgICAgICAgInBsYXRmb3JtIjogewogICAgICAgICAgICAiYXdzIjogewogICAgICAgICAgICAgICAgInJvb3RWb2x1bWUiOiB7CiAgICAgICAgICAgICAgICAgICAgImlvcHMiOiAwLAogICAgICAgICAgICAgICAgICAgICJzaXplIjogMzAwLAogICAgICAgICAgICAgICAgICAgICJ0eXBlIjogImdwMyIKICAgICAgICAgICAgICAgIH0sCiAgICAgICAgICAgICAgICAidHlwZSI6ICJtNS54bGFyZ2UiLAogICAgICAgICAgICAgICAgInpvbmVzIjogWwogICAgICAgICAgICAgICAgICAgICJ1cy1lYXN0LTFhIgogICAgICAgICAgICAgICAgXQogICAgICAgICAgICB9CiAgICAgICAgfSwKICAgICAgICAicmVwbGljYXMiOiAyCiAgICB9LAogICAgInN0YXR1cyI6IHsKICAgICAgICAiY29uZGl0aW9ucyI6IFsKICAgICAgICBdCiAgICB9Cn0K"
            }
        },
    )
    print(response)


# x-ms-original-file: specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/examples/MachinePools_CreateOrUpdate.json
if __name__ == "__main__":
    main()
