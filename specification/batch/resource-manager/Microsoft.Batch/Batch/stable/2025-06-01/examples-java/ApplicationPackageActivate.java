
import com.azure.resourcemanager.batch.models.ActivateApplicationPackageParameters;

/**
 * Samples for ApplicationPackage Activate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/ApplicationPackageActivate.json
     */
    /**
     * Sample code: ApplicationPackageActivate.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void applicationPackageActivate(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.applicationPackages().activateWithResponse("default-azurebatch-japaneast", "sampleacct", "app1", "1",
            new ActivateApplicationPackageParameters().withFormat("zip"), com.azure.core.util.Context.NONE);
    }
}
