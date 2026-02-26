
import com.azure.resourcemanager.postgresqlflexibleserver.models.Configuration;

/**
 * Samples for Configurations Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ConfigurationsUpdate.json
     */
    /**
     * Sample code: Update the value assigned to a specific modifiable configuration (also known as server parameter) of
     * a server.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void updateTheValueAssignedToASpecificModifiableConfigurationAlsoKnownAsServerParameterOfAServer(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        Configuration resource = manager.configurations().getWithResponse("exampleresourcegroup", "exampleserver",
            "constraint_exclusion", com.azure.core.util.Context.NONE).getValue();
        resource.update().withValue("on").withSource("user-override").apply();
    }
}
