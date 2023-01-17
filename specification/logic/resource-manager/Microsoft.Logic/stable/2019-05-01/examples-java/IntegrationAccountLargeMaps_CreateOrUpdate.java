import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.logic.models.MapType;
import java.io.IOException;

/** Samples for IntegrationAccountMaps CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountLargeMaps_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a map larger than 4 MB.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void createOrUpdateAMapLargerThan4MB(com.azure.resourcemanager.logic.LogicManager manager)
        throws IOException {
        manager
            .integrationAccountMaps()
            .define("testMap")
            .withRegion("westus")
            .withExistingIntegrationAccount("testResourceGroup", "testIntegrationAccount")
            .withMapType(MapType.XSLT)
            .withContentType("application/xml")
            .withMetadata(
                SerializerFactory
                    .createDefaultManagementSerializerAdapter()
                    .deserialize("{}", Object.class, SerializerEncoding.JSON))
            .create();
    }
}
