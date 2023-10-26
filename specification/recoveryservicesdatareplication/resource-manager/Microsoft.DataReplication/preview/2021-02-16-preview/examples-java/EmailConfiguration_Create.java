import com.azure.resourcemanager.recoveryservicesdatareplication.models.EmailConfigurationModelProperties;
import java.util.Arrays;

/** Samples for EmailConfiguration Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/EmailConfiguration_Create.json
     */
    /**
     * Sample code: EmailConfiguration_Create.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void emailConfigurationCreate(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .emailConfigurations()
            .define("0")
            .withExistingReplicationVault("rgrecoveryservicesdatareplication", "4")
            .withProperties(
                new EmailConfigurationModelProperties()
                    .withSendToOwners(true)
                    .withCustomEmailAddresses(Arrays.asList("ketvbducyailcny"))
                    .withLocale("vpnjxjvdqtebnucyxiyrjiko"))
            .create();
    }
}
