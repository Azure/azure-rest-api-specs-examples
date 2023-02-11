import com.azure.resourcemanager.costmanagement.models.ExternalCloudProviderType;

/** Samples for Dimensions ByExternalCloudProviderType. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExternalBillingAccountsDimensions.json
     */
    /**
     * Sample code: ExternalBillingAccountDimensionList.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void externalBillingAccountDimensionList(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .byExternalCloudProviderType(
                ExternalCloudProviderType.EXTERNAL_BILLING_ACCOUNTS,
                "100",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
