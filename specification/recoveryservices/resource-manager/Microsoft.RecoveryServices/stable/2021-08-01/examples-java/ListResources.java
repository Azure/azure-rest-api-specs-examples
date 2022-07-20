import com.azure.core.util.Context;

/** Samples for Vaults ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2021-08-01/examples/ListResources.json
     */
    /**
     * Sample code: List of Recovery Services Resources in ResourceGroup.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void listOfRecoveryServicesResourcesInResourceGroup(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.vaults().listByResourceGroup("Default-RecoveryServices-ResourceGroup", Context.NONE);
    }
}
