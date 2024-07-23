
import com.azure.resourcemanager.dataprotection.models.ResourceGuardProxyBase;

/**
 * Samples for DppResourceGuardProxy CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * ResourceGuardProxyCRUD/PutResourceGuardProxy.json
     */
    /**
     * Sample code: Create ResourceGuardProxy.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        createResourceGuardProxy(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.dppResourceGuardProxies().define("swaggerExample")
            .withExistingBackupVault("SampleResourceGroup", "sampleVault")
            .withProperties(new ResourceGuardProxyBase().withResourceGuardResourceId(
                "/subscriptions/f9e67185-f313-4e79-aa71-6458d429369d/resourceGroups/ResourceGuardSecurityAdminRG/providers/Microsoft.DataProtection/resourceGuards/ResourceGuardTestResource"))
            .create();
    }
}
