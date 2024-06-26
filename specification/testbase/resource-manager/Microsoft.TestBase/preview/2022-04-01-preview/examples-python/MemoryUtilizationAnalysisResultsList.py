from azure.identity import DefaultAzureCredential
from azure.mgmt.testbase import TestBase

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-testbase
# USAGE
    python memory_utilization_analysis_results_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = TestBase(
        credential=DefaultAzureCredential(),
        subscription_id="subscription-id",
    )

    response = client.analysis_results.list(
        resource_group_name="contoso-rg1",
        test_base_account_name="contoso-testBaseAccount1",
        package_name="contoso-package2",
        test_result_name="Windows-10-1909-Test-Id",
        analysis_result_type="MemoryUtilization",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/testbase/resource-manager/Microsoft.TestBase/preview/2022-04-01-preview/examples/MemoryUtilizationAnalysisResultsList.json
if __name__ == "__main__":
    main()
