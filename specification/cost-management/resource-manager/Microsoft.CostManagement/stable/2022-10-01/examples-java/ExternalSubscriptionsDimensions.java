import com.azure.resourcemanager.costmanagement.models.ExternalCloudProviderType;

/** Samples for Dimensions ByExternalCloudProviderType. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExternalSubscriptionsDimensions.json
     */
    /**
     * Sample code: ExternalSubscriptionDimensionList.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void externalSubscriptionDimensionList(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .dimensions()
            .byExternalCloudProviderType(
                ExternalCloudProviderType.EXTERNAL_SUBSCRIPTIONS,
                "100",
                null,
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
