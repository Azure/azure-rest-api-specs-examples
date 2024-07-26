
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import java.io.IOException;

/**
 * Samples for LogicApps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/LogicApps_Create.json
     */
    /**
     * Sample code: Create logic app extension.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createLogicAppExtension(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager)
        throws IOException {
        manager.logicApps().define("testcontainerApp0").withExistingContainerApp("examplerg", "testcontainerApp0")
            .withProperties(SerializerFactory.createDefaultManagementSerializerAdapter().deserialize("{}", Object.class,
                SerializerEncoding.JSON))
            .create();
    }
}
