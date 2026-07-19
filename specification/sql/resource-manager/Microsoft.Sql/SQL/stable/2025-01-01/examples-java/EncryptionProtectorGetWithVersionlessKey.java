
import com.azure.resourcemanager.sql.models.EncryptionProtectorName;

/**
 * Samples for EncryptionProtectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/EncryptionProtectorGetWithVersionlessKey.json
     */
    /**
     * Sample code: Get the encryption protector with versionless key.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getTheEncryptionProtectorWithVersionlessKey(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getEncryptionProtectors().getWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645",
            EncryptionProtectorName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
