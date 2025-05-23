
import com.azure.resourcemanager.appcontainers.models.ContainerResources;
import com.azure.resourcemanager.appcontainers.models.JobExecutionContainer;
import com.azure.resourcemanager.appcontainers.models.JobExecutionTemplate;
import java.util.Arrays;

/**
 * Samples for Jobs Start.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/Job_Start.json
     */
    /**
     * Sample code: Run a Container Apps Job.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void runAContainerAppsJob(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.jobs().start("rg", "testcontainerappsjob0",
            new JobExecutionTemplate()
                .withContainers(Arrays.asList(new JobExecutionContainer().withImage("repo/testcontainerappsjob0:v4")
                    .withName("testcontainerappsjob0")
                    .withResources(new ContainerResources().withCpu(0.5D).withMemory("1Gi"))))
                .withInitContainers(Arrays.asList(new JobExecutionContainer().withImage("repo/testcontainerappsjob0:v4")
                    .withName("testinitcontainerAppsJob0").withCommand(Arrays.asList("/bin/sh"))
                    .withArgs(Arrays.asList("-c", "while true; do echo hello; sleep 10;done"))
                    .withResources(new ContainerResources().withCpu(0.5D).withMemory("1Gi")))),
            com.azure.core.util.Context.NONE);
    }
}
