
import com.azure.resourcemanager.durabletask.fluent.models.TransparentDataEncryptionInner;
import com.azure.resourcemanager.durabletask.models.TransparentDataEncryptionKeySource;
import com.azure.resourcemanager.durabletask.models.TransparentDataEncryptionProperties;

/**
 * Samples for TransparentDataEncryptions CreateOrReplace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/TransparentDataEncryptions_CreateOrReplace_MaximumSet_Gen.json
     */
    /**
     * Sample code: TransparentDataEncryptions_CreateOrReplace_MaximumSet.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void transparentDataEncryptionsCreateOrReplaceMaximumSet(
        com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.transparentDataEncryptions().createOrReplace("rgdurabletask", "testscheduler",
            new TransparentDataEncryptionInner().withProperties(new TransparentDataEncryptionProperties()
                .withKeySource(TransparentDataEncryptionKeySource.CUSTOMER_MANAGED)
                .withKeyVaultKeyUri("fakeTokenPlaceholder").withAutoRotationEnabled(true)),
            com.azure.core.util.Context.NONE);
    }
}
