
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/**
 * Samples for Snapshots Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-03-01/examples/Snapshots_Update.json
     */
    /**
     * Sample code: Snapshots_Update.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void snapshotsUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) throws IOException {
        manager.snapshots()
            .update("myRG", "account1", "pool1", "volume1", "snapshot1", SerializerFactory
                .createDefaultManagementSerializerAdapter().deserialize("{}", Object.class, SerializerEncoding.JSON),
                com.azure.core.util.Context.NONE);
    }
}
