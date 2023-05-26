/** Samples for ChildAvailabilityStatuses List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/ChildAvailabilityStatuses_List.json
     */
    /**
     * Sample code: GetChildHealthHistoryByResource.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void getChildHealthHistoryByResource(
        com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager
            .childAvailabilityStatuses()
            .list(
                "subscriptions/227b734f-e14f-4de6-b7fc-3190c21e69f6/resourceGroups/JUHACKETRHCTEST/providers/Microsoft.Compute/virtualMachineScaleSets/rhctest/virtualMachines/4",
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
