/** Samples for AutoScaleVCores Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/deleteAutoScaleVCore.json
     */
    /**
     * Sample code: Delete an auto scale v-core.
     *
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void deleteAnAutoScaleVCore(
        com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        manager
            .autoScaleVCores()
            .deleteByResourceGroupWithResponse("TestRG", "testvcore", com.azure.core.util.Context.NONE);
    }
}
