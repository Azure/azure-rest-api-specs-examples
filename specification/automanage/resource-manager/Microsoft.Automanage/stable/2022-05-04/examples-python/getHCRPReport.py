from azure.identity import DefaultAzureCredential
from azure.mgmt.automanage import AutomanageClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automanage
# USAGE
    python get_hcrp_report.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AutomanageClient(
        credential=DefaultAzureCredential(),
        subscription_id="mySubscriptionId",
    )

    response = client.hcrp_reports.get(
        resource_group_name="myResourceGroupName",
        machine_name="myMachineName",
        configuration_profile_assignment_name="default",
        report_name="b4e9ee6b-1717-4ff0-a8d2-e6d72c33d5f4",
    )
    print(response)


# x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getHCRPReport.json
if __name__ == "__main__":
    main()
