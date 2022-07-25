import com.azure.core.util.Context;

/** Samples for Vaults GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-04-01/examples/GETVault.json
     */
    /**
     * Sample code: Get Recovery Services Resource.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void getRecoveryServicesResource(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager
            .vaults()
            .getByResourceGroupWithResponse("Default-RecoveryServices-ResourceGroup", "swaggerExample", Context.NONE);
    }
}
