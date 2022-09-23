import com.azure.resourcemanager.security.models.AwsEnvironmentData;
import com.azure.resourcemanager.security.models.CloudName;
import com.azure.resourcemanager.security.models.CspmMonitorAwsOffering;
import com.azure.resourcemanager.security.models.CspmMonitorAwsOfferingNativeCloudConnection;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for SecurityConnectors CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-08-01-preview/examples/SecurityConnectors/PutSecurityConnector_example.json
     */
    /**
     * Sample code: Create or update a security connector.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void createOrUpdateASecurityConnector(com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .securityConnectors()
            .define("exampleSecurityConnectorName")
            .withRegion("Central US")
            .withExistingResourceGroup("exampleResourceGroup")
            .withTags(mapOf())
            .withHierarchyIdentifier("exampleHierarchyId")
            .withEnvironmentName(CloudName.AWS)
            .withOfferings(
                Arrays
                    .asList(
                        new CspmMonitorAwsOffering()
                            .withNativeCloudConnection(
                                new CspmMonitorAwsOfferingNativeCloudConnection()
                                    .withCloudRoleArn("arn:aws:iam::00000000:role/ASCMonitor"))))
            .withEnvironmentData(new AwsEnvironmentData())
            .create();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
