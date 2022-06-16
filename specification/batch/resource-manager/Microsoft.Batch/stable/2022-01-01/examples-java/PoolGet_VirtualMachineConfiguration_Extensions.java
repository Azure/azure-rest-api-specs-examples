import com.azure.core.util.Context;

/** Samples for Pool Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PoolGet_VirtualMachineConfiguration_Extensions.json
     */
    /**
     * Sample code: GetPool - VirtualMachineConfiguration Extensions.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void getPoolVirtualMachineConfigurationExtensions(
        com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().getWithResponse("default-azurebatch-japaneast", "sampleacct", "testpool", Context.NONE);
    }
}
