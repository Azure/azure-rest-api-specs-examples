import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.logic.models.AssemblyProperties;
import java.io.IOException;

/** Samples for IntegrationAccountAssemblies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAssemblies_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update an account assembly.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void createOrUpdateAnAccountAssembly(com.azure.resourcemanager.logic.LogicManager manager)
        throws IOException {
        manager
            .integrationAccountAssemblies()
            .define("testAssembly")
            .withRegion("westus")
            .withExistingIntegrationAccount("testResourceGroup", "testIntegrationAccount")
            .withProperties(
                new AssemblyProperties()
                    .withMetadata(
                        SerializerFactory
                            .createDefaultManagementSerializerAdapter()
                            .deserialize("{}", Object.class, SerializerEncoding.JSON))
                    .withContent("Base64 encoded Assembly Content")
                    .withAssemblyName("System.IdentityModel.Tokens.Jwt"))
            .create();
    }
}
