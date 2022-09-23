import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.security.fluent.models.ApplicationInner;
import com.azure.resourcemanager.security.models.ApplicationSourceResourceType;
import java.io.IOException;
import java.util.Arrays;

/** Samples for SecurityConnectorApplicationOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/PutSecurityConnectorApplication_example.json
     */
    /**
     * Sample code: Create Application.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void createApplication(com.azure.resourcemanager.security.SecurityManager manager)
        throws IOException {
        manager
            .securityConnectorApplicationOperations()
            .createOrUpdateWithResponse(
                "gcpResourceGroup",
                "gcpconnector",
                "ad9a8e26-29d9-4829-bb30-e597a58cdbb8",
                new ApplicationInner()
                    .withDisplayName("GCP Admin's application")
                    .withDescription("An application on critical GCP recommendations")
                    .withSourceResourceType(ApplicationSourceResourceType.ASSESSMENTS)
                    .withConditionSets(
                        Arrays
                            .asList(
                                SerializerFactory
                                    .createDefaultManagementSerializerAdapter()
                                    .deserialize(
                                        "{\"conditions\":[{\"operator\":\"contains\",\"property\":\"$.Id\",\"value\":\"-prod-\"}]}",
                                        Object.class,
                                        SerializerEncoding.JSON))),
                Context.NONE);
    }
}
