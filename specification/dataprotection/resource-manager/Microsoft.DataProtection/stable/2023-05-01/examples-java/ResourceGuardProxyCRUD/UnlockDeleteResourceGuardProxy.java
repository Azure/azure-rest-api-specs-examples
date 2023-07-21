import com.azure.resourcemanager.dataprotection.models.UnlockDeleteRequest;
import java.util.Arrays;

/** Samples for DppResourceGuardProxy UnlockDelete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/ResourceGuardProxyCRUD/UnlockDeleteResourceGuardProxy.json
     */
    /**
     * Sample code: UnlockDelete ResourceGuardProxy.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void unlockDeleteResourceGuardProxy(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .dppResourceGuardProxies()
            .unlockDeleteWithResponse(
                "SampleResourceGroup",
                "sampleVault",
                "swaggerExample",
                new UnlockDeleteRequest()
                    .withResourceGuardOperationRequests(
                        Arrays
                            .asList(
                                "/subscriptions/f9e67185-f313-4e79-aa71-6458d429369d/resourceGroups/ResourceGuardSecurityAdminRG/providers/Microsoft.DataProtection/resourceGuards/ResourceGuardTestResource/deleteBackupInstanceRequests/default"))
                    .withResourceToBeDeleted(
                        "/subscriptions/5e13b949-1218-4d18-8b99-7e12155ec4f7/resourceGroups/SampleResourceGroup/providers/Microsoft.DataProtection/backupVaults/sampleVault/backupInstances/TestBI9779f4de"),
                com.azure.core.util.Context.NONE);
    }
}
