import com.azure.core.util.Context;

/** Samples for Usages ListByVaults. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-04-01/examples/ListUsages.json
     */
    /**
     * Sample code: Gets vault usages.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void getsVaultUsages(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.usages().listByVaults("Default-RecoveryServices-ResourceGroup", "swaggerExample", Context.NONE);
    }
}
